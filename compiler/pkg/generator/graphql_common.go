package generator

import (
	"fmt"
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	log "github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/graph-framework-for-microservices/compiler/pkg/parser"
	"github.com/vmware-tanzu/graph-framework-for-microservices/compiler/pkg/util"
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

const (
	//GraphQL standard data types
	Int     string = "Int"
	Float   string = "Float"
	String  string = "String"
	Boolean string = "Boolean"

	// Custom query
	CustomQuerySchema = `Id: ID
	ParentLabels: Map
`
)

type ReturnStatement struct {
	Alias       string
	ReturnType  string
	FieldCount  int
	CRDName     string
	ChainAPI    string
	IsSingleton bool
}

type FieldProperty struct {
	IsResolver              bool
	IsNexusTypeField        bool
	IsNexusOrSingletonField bool
	IsChildOrLink           bool
	IsChildrenOrLinks       bool
	IsMapTypeField          bool
	IsArrayTypeField        bool
	IsStdTypeField          bool
	IsCustomTypeField       bool
	IsPointerTypeField      bool
	IsStringType            bool
	IsArrayStdType          bool
	IsSingleton             bool
	PkgName                 string
	NodeName                string
	FieldName               string
	FieldType               string
	FieldTypePkgPath        string
	ModelType               string
	SchemaFieldName         string
	SchemaTypeName          string
	BaseTypeName            string
	Alias                   string
	ReturnType              string
	FieldCount              int
	CRDName                 string
	ChainAPI                string
	LinkAPI                 string
}

type NodeProperty struct {
	IsParentNode           bool
	HasParent              bool
	IsSingletonNode        bool
	IsNexusNode            bool
	BaseImportPath         string
	CrdName                string
	ResolverCount          int
	PkgName                string
	NodeName               string
	SchemaName             string
	Alias                  string
	ReturnType             string
	GroupResourceNameTitle string
	ChildFields            []FieldProperty
	LinkFields             []FieldProperty
	ChildrenFields         []FieldProperty
	LinksFields            []FieldProperty
	ArrayFields            []FieldProperty
	CustomFields           []FieldProperty
	NonStructFields        []FieldProperty
	GraphqlSchemaFields    []FieldProperty
	ResolverFields         map[string][]FieldProperty
	CustomQueries          []nexus.GraphQLQuery
	GraphQlSpec            nexus.GraphQLSpec
}

// Convert go standardType to GraphQL standardType
func convertGraphqlStdType(t string) string {
	// remove pointers
	typeWithoutPointer := strings.ReplaceAll(t, "*", "")
	switch typeWithoutPointer {
	case "string":
		return String
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return Int
	case "bool":
		return Boolean
	case "float32", "float64":
		return Float
	default:
		return ""
	}
}

func jsonMarshalResolver(FieldName, PkgName string) string {
	return fmt.Sprintf("%s, _ := json.Marshal(v%s.Spec.%s)\n%sData := string(%s)\n", FieldName, PkgName, FieldName, FieldName, FieldName)
}

func getPkgName(pkgs parser.Packages, pkgPath string) string {
	importPath, err := strconv.Unquote(pkgPath)
	if err != nil {
		log.Errorf("Failed to parse the package path : %s: %v", pkgPath, err)
	}
	return pkgs[importPath].Name
}

func genSchemaResolverName(fn1, fn2 string) (string, string) {
	return fmt.Sprintf("%s_%s", strings.ToLower(fn1), fn2), cases.Title(language.Und, cases.NoLower).String(fn1) + cases.Title(language.Und, cases.NoLower).String(fn2)
}

func ValidateImportPkg(pkgName, typeString string, importMap map[string]string, pkgs parser.Packages) (string, string) {
	typeWithoutPointers := strings.ReplaceAll(typeString, "*", "")
	if strings.Contains(typeWithoutPointers, ".") {
		part := strings.Split(typeWithoutPointers, ".")
		if val, ok := importMap[part[0]]; ok {
			pkgName := getPkgName(pkgs, val)
			repName := strings.ReplaceAll(pkgName, "-", "")
			return genSchemaResolverName(repName, part[1])
		}
		for _, v := range importMap {
			pkgName := getPkgName(pkgs, v)
			if pkgName == part[0] {
				repName := strings.ReplaceAll(pkgName, "-", "")
				return genSchemaResolverName(repName, part[1])
			}
		}
		return genSchemaResolverName(pkgName, part[1])
	}
	return genSchemaResolverName(pkgName, typeWithoutPointers)
}

func getBaseNodeType(typeString string) string {
	if strings.Contains(typeString, ".") {
		part := strings.Split(typeString, ".")
		return part[1]
	}
	return typeString
}

func constructNexusTypeMap(nodes []*NodeProperty) map[string]string {
	crdNameMap := make(map[string]string)
	for _, n := range nodes {
		if n.IsNexusNode || n.IsSingletonNode {
			crdNameMap[n.CrdName] = n.PkgName + n.NodeName
		}
	}
	return crdNameMap
}

func findTypeAndPkgForField(ptParts []string, importMap map[string]string, pkgs map[string]parser.Package) (string, *parser.Package) {
	structPkg := ptParts[0]
	structType := ptParts[1]

	pkgPath, ok := importMap[structPkg]
	if !ok {
		log.Errorf("Cannot find the package name %s for the type %s", structPkg, structType)
		return "", nil
	}

	importPath, err := strconv.Unquote(pkgPath)
	if err != nil {
		log.Errorf("Failed to parse package %s for the type %s with error %v", pkgPath, structType, err)
		return "", nil
	}

	p, ok := pkgs[importPath]
	if !ok {
		log.Errorf("Cannot find the package details from the path %s for the type %s", importPath, structType)
		return "", nil
	}

	return structType, &p
}

/*
collect and construct type alias field into map recursively before
populating the nexus node and custom struct type
ex. nonStructMap[pkgName] = nodeType  -->  nonStructMap["root"] = "AliasTypeFoo"
*/
func constructAliasType(sortedPackages []parser.Package) map[string]string {
	nonStructMap := make(map[string]string)
	for _, pkg := range sortedPackages {
		for _, node := range pkg.GetNonStructTypes() {
			pkgName := fmt.Sprintf("%s_%s", pkg.Name, parser.GetTypeName(node))
			// NonStruct Map
			nonStructType := types.ExprString(node.Type)
			nonStructMap[pkgName] = nonStructType
		}
	}

	return nonStructMap
}

func setNexusProperties(nodeHelper parser.NodeHelper, node *ast.TypeSpec, nodeProp *NodeProperty) {
	if len(nodeHelper.Parents) > 0 {
		nodeProp.HasParent = true
	}

	if parser.IsSingletonNode(node) {
		nodeProp.IsSingletonNode = true
	}

	if parser.IsNexusNode(node) {
		nodeProp.IsNexusNode = true
	}
}

// isRootOfGraph intended to allow only one root of the graph,
// if we receive multiple node in such behaviour, then we allow the first node and the rest will be skipped with error
// arg: `parents` indicates the node's parents, `rootOfGraph` indicates if the received node is the root of the graph or not.
func isRootOfGraph(parents []string, rootOfGraph bool) bool {
	if len(parents) == 0 && !rootOfGraph {
		return true
	}

	return rootOfGraph
}

func getGraphqlSchemaName(pattern, fieldName, schemaType string) string {
	if fieldName != "" {
		return fmt.Sprintf(pattern, util.GetTag(fieldName), schemaType)
	}
	return fmt.Sprintf(pattern, fieldName, schemaType)
}

func getTsmGraphqlSchemaFieldName(sType GraphQLSchemaType, fieldName, schemaType string, f *ast.Field) string {
	pattern := ""
	nonNullable := parser.IsNexusGraphqlNonNullField(f)
	switch sType {
	case Standard, JsonMarshal, Child, Link:
		if nonNullable {
			pattern = "%s: %s!"
		} else {
			pattern = "%s: %s"
		}
	case NamedChild:
		pattern = "%s(Id: ID): [%s!]"
	}
	schemaName := getGraphqlSchemaName(pattern, fieldName, schemaType)
	if parser.IsTsmGraphqlDirectivesField(f) {
		replacer := strings.NewReplacer("nexus-graphql-tsm-directive:", "", "\\", "")
		out := replacer.Replace(parser.GetTsmGraphqlDirectives(f))
		fmt.Println("getDirectives:", schemaName, strings.Trim(out, "\""))
		schemaName += " " + strings.Trim(out, "\"")
	}
	return schemaName
}

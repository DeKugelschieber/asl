package asl

import (
    "strconv"
)

const TAB = "    "

func Parse(token []Token) string {
    initParser(token)
    
    for tokenIndex < len(token) {
        parseBlock()
    }
    
    return out
}

func parseBlock() {
    if accept("var") {
        parseVar()
    } else if accept("if") {
        parseIf()
    } else if accept("func") {
        parseFunction()
    } else {
        parseStatement()
    }
}

func parseVar() {
    expect("var")
    appendOut(get().token)
    next()
    
    if accept("=") {
        next()
        appendOut(" = "+get().token)
        next()
    }
    
    appendOut(";\n")
    expect(";")
}

func parseIf() {
    expect("if")
    appendOut("if (")
    parseCondition()
    appendOut(") then {\n")
    expect("{")
    parseBlock()
    expect("}")
    
    if accept("else") {
        next()
        expect("{")
        appendOut("} else {\n")
        parseBlock()
        expect("}")
    }
    
    appendOut("};\n")
}

func parseCondition() {
    for !accept("{") {
        appendOut(get().token)
        next()
        
        if !accept("{") {
            appendOut(" ")
        }
    }
}

func parseFunction() {
    expect("func")
    appendOut(get().token+" = {\n")
    next()
    expect("(")
    parseFunctionParameter()
    expect(")")
    expect("{")
    parseBlock()
    expect("}")
    appendOut("};\n")
}

func parseFunctionParameter() {
    // empty parameter list
    if accept(")") {
        return;
    }
    
    i := int64(0)
    
    for !accept(")") {
        name := get().token
        next()
        appendOut(name+" = this select "+strconv.FormatInt(i, 10)+";\n")
        i++
        
        if !accept(")") {
            expect(",")
        }
    }
}

// Everything that does not start with a keyword.
func parseStatement() {
    // empty block
    if accept("}") {
        return;
    }
    
    // variable or function name
    name := get().token
    next()
    
    if accept("=") {
        appendOut(name)
        parseAssignment()
    } else {
        parseFunctionCall()
        appendOut(name+";\n")
    }
}

func parseAssignment() {
    expect("=")
    appendOut(" = "+get().token)
    next()
    expect(";")
    appendOut(";\n")
}

func parseFunctionCall() {
    expect("(")
    params := parseParameter()
    expect(")")
    expect(";")
    appendOut("["+params+"] call ")
}

func parseParameter() string {
    params := ""
    
    for !accept(")") {
        params += get().token
        next()
        
        if !accept(")") {
            expect(",")
            params += ", "
        }
    }
    
    return params
}

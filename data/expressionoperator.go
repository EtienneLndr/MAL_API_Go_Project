/**
 * MIT License
 *
 * Copyright (c) 2018 CNES
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package data

import (
	. "github.com/ccsdsmo/malgo/mal"
)

// ExpressionOperator holds a set of possible expression operators
type ExpressionOperator UOctet

const (
	COM_EXPRESSIONOPERATOR_EQUAL ExpressionOperator = iota + 1
	COM_EXPRESSIONOPERATOR_DIFFER
	COM_EXPRESSIONOPERATOR_GREATER
	COM_EXPRESSIONOPERATOR_GREATER_OR_EQUAL
	COM_EXPRESSIONOPERATOR_LESS
	COM_EXPRESSIONOPERATOR_LESS_OR_EQUAL
	COM_EXPRESSIONOPERATOR_CONTAINS
	COM_EXPRESSIONOPERATOR_ICONTAINS
)

const (
	COM_EXPRESSIONOPERATOR_TYPE_SHORT_FORM Integer = 0x05
	COM_EXPRESSIONOPERATOR_SHORT_FORM      Long    = 0x2000201000005
)

// TransformOperator transforms an ExpressionOperator to a String
func (e ExpressionOperator) TransformOperator() String {
	switch e {
	case COM_EXPRESSIONOPERATOR_EQUAL:
		return "="
	case COM_EXPRESSIONOPERATOR_DIFFER:
		return "!="
	case COM_EXPRESSIONOPERATOR_GREATER:
		return ">"
	case COM_EXPRESSIONOPERATOR_GREATER_OR_EQUAL:
		return ">="
	case COM_EXPRESSIONOPERATOR_LESS:
		return "<"
	case COM_EXPRESSIONOPERATOR_LESS_OR_EQUAL:
		return "<="
	case COM_EXPRESSIONOPERATOR_CONTAINS:
		return "LIKE '%"
	case COM_EXPRESSIONOPERATOR_ICONTAINS:
		return "NOT LIKE '%"
	default:
		return ""
	}
}

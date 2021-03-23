package nexe

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

func TranspileBinary(expr *ast.Binary, scope *Scope) (Stmt, []*types.Type, error) {
	left, leftType, err := TranspileStmt(expr.Left, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	right, _, err := TranspileStmt(expr.Right, scope)
	if err != nil {
		return Expr(""), nil, err
	}

	switch expr.Op {
	case lexer.TokenPlusAssign:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("%s = %s.add(%s)", left, left, right)), leftType, nil
		}

		return Expr(fmt.Sprintf("%s += (%s)", left, right)), leftType, nil

	case lexer.TokenPlus:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("(%s).add(%s)", left, right)), leftType, nil
		}

		return Expr(fmt.Sprintf("(%s) + (%s)", left, right)), leftType, nil

	case lexer.TokenMinus:
		return Expr(fmt.Sprintf("(%s).minus(%s)", left, right)), leftType, nil

	case lexer.TokenMinusAssign:
		return Expr(fmt.Sprintf("%s = %s.minus(%s)", left, left, right)), leftType, nil

	case lexer.TokenTimes:
		return Expr(fmt.Sprintf("(%s).mul(%s)", left, right)), leftType, nil

	case lexer.TokenTimesAssign:
		return Expr(fmt.Sprintf("%s = %s.mul(%s)", left, left, right)), leftType, nil

	case lexer.TokenDivide:
		return Expr(fmt.Sprintf("(%s).div(%s)", left, right)), leftType, nil

	case lexer.TokenDivideAssign:
		return Expr(fmt.Sprintf("%s = %s.div(%s)", left, left, right)), leftType, nil

	case lexer.TokenRemainder:
		return Expr(fmt.Sprintf("(%s).mod(%s)", left, right)), leftType, nil

	case lexer.TokenRemainderAssign:
		return Expr(fmt.Sprintf("%s = %s.mod(%s)", left, left, right)), leftType, nil

	case lexer.TokenGreaterThan:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("(%s).gt(%s)", left, right)), []*types.Type{types.Bool}, nil
		}

		return Expr(fmt.Sprintf("(%s) > (%s)", left, right)), []*types.Type{types.Bool}, nil

	case lexer.TokenGreaterThanEqual:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("(%s).gte(%s)", left, right)), []*types.Type{types.Bool}, nil
		}

		return Expr(fmt.Sprintf("(%s) >= (%s)", left, right)), []*types.Type{types.Bool}, nil

	case lexer.TokenLessThan:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("(%s).lt(%s)", left, right)), []*types.Type{types.Bool}, nil
		}

		return Expr(fmt.Sprintf("(%s) < (%s)", left, right)), []*types.Type{types.Bool}, nil

	case lexer.TokenLessThanEqual:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("(%s).lte(%s)", left, right)), []*types.Type{types.Bool}, nil
		}

		return Expr(fmt.Sprintf("(%s) <= (%s)", left, right)), []*types.Type{types.Bool}, nil

	case lexer.TokenEqual:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("(%s).eq(%s)", left, right)), []*types.Type{types.Bool}, nil
		}

		return Expr(fmt.Sprintf("(%s) === (%s)", left, right)), []*types.Type{types.Bool}, nil

	case lexer.TokenNotEqual:
		if leftType[0].Kind == types.KindNumber {
			return Expr(fmt.Sprintf("!(%s).eq(%s)", left, right)), []*types.Type{types.Bool}, nil
		}

		return Expr(fmt.Sprintf("(%s) !== (%s)", left, right)), []*types.Type{types.Bool}, nil
	}

	panic(expr.Op)
}

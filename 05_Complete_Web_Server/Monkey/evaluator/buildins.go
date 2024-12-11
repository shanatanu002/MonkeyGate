package evaluator

import (
	"interpreter/Monkey/object"
)

var buildins = map[string]*object.Buildin{
	"len": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Hash:
				return &object.Integer{Value: int64(len(arg.Pairs))}
			default:
				return newError("argument to `len` is not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be array, got = %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return newError("array is empty")
		},
	},
	"last": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be array, got = %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return newError("array is empty")
		},
	},
	"rest": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be array, got = %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return newError("array is empty : cannot find `rest`")
		},
	},
	"push": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got = %d, want = 2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument_1 to `push` must be array, got = %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"part": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got = %d, want = 3", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument_1 to `part` must be array, got = %s", args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("argument_2 to `part` must be integer, got = %s", args[1].Type())
			}
			if args[2].Type() != object.INTEGER_OBJ {
				return newError("argument_3 to `part` must be integer, got = %s", args[2].Type())
			}

			arr := args[0].(*object.Array)
			start := args[1].(*object.Integer).Value
			end := args[2].(*object.Integer).Value
			length := len(arr.Elements)
			if start > end {
				return newError("argument_2 cannot be greater than argument_3")
			}

			if start < 0 || start >= int64(length) || end < 0 || end > int64(length) {
				return newError("index out of bounds. Check argument_2 and argument_3")
			}
			partElements := make([]object.Object, end-start)
			copy(partElements, arr.Elements[start:end])

			return &object.Array{Elements: partElements}
		},
	},
	"puts": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			var output string
			for _, arg := range args {
				output += arg.Inspect() + "\n"
			}
			// instead of printing, return the captured output
			return &object.String{Value: output}
		},
	},
}

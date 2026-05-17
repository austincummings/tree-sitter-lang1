; Indent the contents of any node delimited by braces, brackets, or parens.
; The `@end` capture marks the closing delimiter so it lines up with the
; opener instead of staying indented.

(_ "{" "}" @end) @indent
(_ "[" "]" @end) @indent
(_ "(" ")" @end) @indent

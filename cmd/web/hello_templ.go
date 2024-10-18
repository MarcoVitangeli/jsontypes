// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func App() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><link href=\"https://cdn.jsdelivr.net/npm/flowbite@2.5.2/dist/flowbite.min.css\" rel=\"stylesheet\"><script src=\"https://cdnjs.cloudflare.com/ajax/libs/ace/1.4.14/ace.js\"></script></head></html><body class=\"p-8\"><h1 class=\"mb-2 text-2xl font-extrabold leading-none tracking-tight text-gray-900 md:text-3xl lg:text-3xl dark:text-white\">Convert JSON objects into Go types</h1><p class=\"mb-3 text-gray-500 dark:text-gray-400\">The conversion works such that a JSON object will be recursively traversed (until a maximum depth) and it will generate a Go file with the types to parse it. The system supports the following types:</p><ul class=\"max-w-md space-y-1 text-gray-500 list-disc list-inside dark:text-gray-400\"><li>float64</li><li>time.Time</li><li>string</li><li>bool</li><li>Nested objects</li><li>Arrays of all previous elements</li></ul><div class=\"flex justify-between gap-6\" style=\"margin-top: 2rem;\"><div class=\"w-1/2\"><label for=\"input-editor\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">JSON Input</label><div id=\"input-editor\" class=\"block p-4 w-full h-96 bg-gray-50 rounded-lg border border-gray-300 dark:bg-gray-700 dark:border-gray-600\"></div></div><div class=\"w-1/2\"><label for=\"output-editor\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Generated Go Code</label><div id=\"output-editor\" class=\"block p-4 w-full h-96 bg-gray-50 rounded-lg border border-gray-300 dark:bg-gray-700 dark:border-gray-600\"></div></div></div><button id=\"generate-btn\" hx-post=\"/api/generate\" hx-swap=\"none\" hx-headers=\"{&#34;Content-Type&#34;: &#34;application/json&#34;}\" hx-vals=\"js:{&#34;json&#34;: JSON.stringify(inputEditor.getValue())}\" class=\"mt-4 px-4 py-2 bg-blue-500 text-white rounded\">Generate Go Code</button><script>\n        var inputEditor = ace.edit(\"input-editor\");\n        inputEditor.setTheme(\"ace/theme/github\");\n        inputEditor.session.setMode(\"ace/mode/json\");\n        inputEditor.setFontSize(16);\n\n        var outputEditor = ace.edit(\"output-editor\");\n        outputEditor.setTheme(\"ace/theme/github\");\n        outputEditor.session.setMode(\"ace/mode/golang\");\n        outputEditor.setFontSize(16);\n        outputEditor.setReadOnly(true);\n\n        document.body.addEventListener('htmx:afterRequest', function(evt) {\n        if (evt.detail.target.id === 'generate-btn') {\n            let goCode = evt.detail.xhr.responseText;\n            outputEditor.setValue(goCode, 1);\n        }\n    });\n    </script><script src=\"https://unpkg.com/htmx.org\"></script><script src=\"https://cdn.jsdelivr.net/npm/flowbite@2.5.2/dist/flowbite.min.js\"></script></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

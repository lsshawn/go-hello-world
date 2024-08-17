// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/lsshawn/go-todo/views/layout"
import "github.com/lsshawn/go-todo/internal/dto"
import "github.com/lsshawn/go-todo/views/components"

func Index(todos []*dto.TodoCardDto) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"h-screen w-screen bg-black flex flex-col justify-center items-center relative\"><div class=\"w-screen bg-black flex flex-col justify-center items-center relative max-w-2xl\" x-data=\"{\n          inputText: &#39;&#39;,\n          isSubmitting: false,\n          resetInput() { console.log(&#39;reset&#39;); this.$refs.todoInputText.focus(); this.inputText = &#39;&#39;; this.isSubmitting = true; },\n          init() {\n            this.$nextTick(() =&gt; {\n                document.body.addEventListener(&#39;htmx:afterOnLoad&#39;, () =&gt; {\n                    this.isSubmitting = false;\n                });\n            });\n        }\n        }\"><form class=\"flex flex-col w-full gap-4\" hx-post=\"/add-todo\" hx-swap=\"beforeend\" hx-target=\"#todos\" hx-indicator=\"#spinner\" @submit.prevent=\"resetInput()\"><label for=\"add-todo-input\">Todo</label> <input type=\"text\" id=\"add-todo-input\" name=\"text\" class=\"rounded-2xl text-black font-sans text-sm flex-1\" x-model=\"inputText\" x-ref=\"todoInputText\"> <button id=\"add-todo-button\" type=\"submit\" class=\"btn btn-success\" :disabled=\"isSubmitting\"><span class=\"loading loading-spinner htmx-indicator\" id=\"spinner\" role=\"status\" aria-hidden=\"true\"></span> Add TODO</button></form><div class=\"h-12\"></div><div class=\"bg-white p-6 rounded-2xl shadow-lg w-full flex flex-col gap-4\" id=\"todos\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, todo := range todos {
				templ_7745c5c3_Err = components.TodoCard(todo).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

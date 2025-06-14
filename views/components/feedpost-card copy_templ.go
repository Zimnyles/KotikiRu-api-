// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "time"
import "github.com/rvflash/elapsed"

type FeedPostCardProps struct {
	CreatorLogin      string
	Content           string
	ImagePath         string
	CreatedAt         time.Time
	CreatorAvatarPath string
}

func FeedPostCard(props FeedPostCardProps) templ.Component {
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
		date := elapsed.LocalTime(props.CreatedAt, "ru")
		templ_7745c5c3_Err = FeedPostCardStyle().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"post-card\"><div class=\"post-card__footer\"><div class=\"post-card__footer--creator\"><img class=\"profile-avatar\" crossorigin=\"anonymous\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.CreatorAvatarPath)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/feedpost-card copy.templ`, Line: 22, Col: 98}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" alt=\"Изображение профиля\"> <span class=\"post-card__creatorlogin\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.CreatorLogin)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/feedpost-card copy.templ`, Line: 23, Col: 77}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</span></div><span class=\"post-card__createdat\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(date)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/feedpost-card copy.templ`, Line: 25, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</span></div><div class=\"post-card__content__text\"><span class=\"post-card__text\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(props.Content)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/feedpost-card copy.templ`, Line: 29, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</span></div><div class=\"post-card__content\"><div class=\"post-card__content__image\"><img class=\"post-image\" crossorigin=\"anonymous\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.ImagePath)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/feedpost-card copy.templ`, Line: 33, Col: 86}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\" alt=\"Изображение поста\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func FeedPostCardStyle() templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<style>\r\n\r\n        .post-card__text{\r\n            white-space: pre-wrap; /* Сохраняет переносы строк из текста */\r\n            word-wrap: break-word; /* Переносит длинные слова */\r\n            overflow-wrap: break-word; /* Альтернатива для word-wrap */\r\n            word-break: break-word; /* Оптимальный перенос слов */\r\n            display: block; /* Если нужно чтобы span вел себя как блок */\r\n            width: 100%; /* Занимает всю доступную ширину */\r\n        }\r\n\r\n        .post-card__content__text{\r\n            margin-left: 15px;\r\n            margin-right: 15px;\r\n        }\r\n\r\n        .post-image{\r\n            border-radius: 20px;\r\n            object-fit: scale-down;\r\n            max-width: 100%;\r\n            max-height: 100%;\r\n            width: auto;\r\n            height: auto;\r\n        }\r\n\r\n        .post-card__content{\r\n           display: flex;\r\n           flex-grow: 1;\r\n           justify-content: center;\r\n        }\r\n\r\n        .post-card__creatorlogin{\r\n            margin-left: 10px;\r\n        }\r\n\r\n        .post-card__footer--creator{\r\n            display: flex;\r\n            flex-direction: row;\r\n            align-items: center;\r\n            justify-content: center;\r\n        }\r\n\r\n        .post-card__content__image{\r\n            object-fit: cover;\r\n            border-radius: 20px;\r\n            max-width: 800px;\r\n            max-height: 450px;\r\n            width: auto;\r\n            height: auto;\r\n            overflow: hidden;\r\n            display: flex;\r\n            justify-content: center;\r\n            align-items: center;\r\n            margin-bottom: 15px;\r\n        }\r\n\r\n        .profile-avatar {\r\n\r\n            width: 30px;\r\n            height: 30px;\r\n            border-radius: 50%;\r\n            object-fit: cover;\r\n            \r\n        }\r\n        .post-card__footer{\r\n            display: flex;\r\n            justify-content: space-between; /* Равномерное распределение */\r\n            padding-right: 15px;\r\n            padding-left: 15px;\r\n            margin: 0 auto; /* Центрирование */\r\n            align-items: center;\r\n            box-sizing: border-box; /* Чтобы padding не увеличивал ширину */\r\n        \r\n            background: #303030;\r\n            max-width: 1100px;\r\n            width: 100%;\r\n            height: 40px;\r\n            border-radius: 20px 20px 0 0\r\n        }\r\n\r\n        .post-card {\r\n            overflow: hidden;\r\n            color: var(--color-white);\r\n            display: flex;\r\n            max-width: 1100px;\r\n            width: 100%;\r\n            min-height: 400px;\r\n            height: 100%;\r\n            flex-direction: column;\r\n            background: #222222;\r\n            border-radius: 20px;\r\n            gap: 15px ;\r\n            width: 100%;\r\n\r\n        \r\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate

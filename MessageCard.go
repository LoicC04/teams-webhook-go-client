package teams_client

const MessageCardTemplate string = `{
  "@context": "https://schema.org/extensions",
  "@type": "MessageCard",
  "correlationId": "{{.CorrelationId}}",
  "summary": "{{.Summary}}",
  "text" : "{{.Text}}",
  "themeColor": "{{.ThemeColor}}",
  "title": "{{.Title}}",
  "sections": [
    {{range $i, $f := .Sections}}
    {
    "title": "{{.Title}}",
    "activityTitle": "{{.Activity.Title}}",
    "activitySubtitle": "{{.Activity.Subtitle}}",
    "activityImage": "{{.Activity.Image}}",
    "activityText": "{{.Activity.Text}}",
    "facts": [
      {{range $i, $f := .Facts}}
      {
        "name": "{{$f.Name}}",
        "value": "{{$f.Value}}"
      },
      {{end}}
    ],
    "potentialAction": [
      {{range $i, $f := .PotentialActions}}
      {
        "@type": "OpenUri",
        "name": "{{.Name}}",
        "targets": [
          {
            "os": "default",
            "uri": "{{.Uri}}"
          },
        ]
      }
      {{end}}
    ]
  }
    {{end}}
  ]
}`
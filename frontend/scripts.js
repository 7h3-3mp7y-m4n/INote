marked.setOptions({
  breaks: true,
  highlight: (code, lang) => {
    return hljs.highlightAuto(code, [lang]).value;
  },
  langPrefix: 'hljs language-',
});

const editor = document.getElementById("editor");
const preview = document.getElementById("preview");

editor.addEventListener("input", () => {
  const raw = editor.value;
  const html = marked.parse(raw);
  preview.innerHTML = html;
});

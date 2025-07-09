marked.setOptions({
  breaks: true,
  highlight: (code, lang) => hljs.highlightAuto(code, [lang]).value,
  langPrefix: 'hljs language-',
});

const editor = document.getElementById("editor");
const preview = document.getElementById("preview");
const saveBtn = document.getElementById("saveBtn");

editor.addEventListener("input", () => {
  preview.innerHTML = marked.parse(editor.value);
});

saveBtn.addEventListener("click", async () => {
  const content = editor.value;
  const format = document.getElementById("format").value;
  const expires = parseInt(document.getElementById("expires").value);
  const expiresAt = new Date(Date.now() + expires * 1000).toISOString();

  const res = await fetch("/notes", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ content, format, expires_at: expiresAt })
  });

  if (!res.ok) {
    const err = await res.json();
    alert("Error saving note: " + err.error);
    return;
  }

  const data = await res.json();
  window.location.href = `/notes/${data.slug}`;
});

const addTargetToMarkdown = () => {
  const markdown = document.getElementById("markdown");
  const markdownAnchors = markdown.querySelectorAll("a");

  markdownAnchors.forEach((anchor, i) => {
    anchor.target = "_blank";
  });
};
addTargetToMarkdown();

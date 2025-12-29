window.addEventListener("load", () => {
  console.log("Page loaded");

  const input = document.getElementById("todo-input");
  const addBtn = document.getElementById("add-button");

  addBtn.addEventListener("click", () => {
    const description = input.value.trim();

    if (description === "") {
      alert("Todo cannot be empty");
      return;
    }

    fetch("/todos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ description })
    })
      .then(res => res.json())
      .then(() => {
        input.value = "";
        return fetch("/todos");
      })
      .then(res => res.json())
      .then(renderTodos)
      .catch(err => console.error(err));
  });

  fetch("/todos")
    .then(res => res.json())
    .then(renderTodos)
    .catch(err => console.error("Fetch error:", err));
});

function renderTodos(todos) {
  const list = document.getElementById("todo-list");
  list.innerHTML = "";

  todos.forEach(todo => {
    const li = document.createElement("li");
    li.textContent = todo.Description;
    list.appendChild(li);
  });
}
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
  const list  = document.getElementById("todo-list");
  list.innerHTML = "";
  
  todos.forEach(todo => {
    const li = document.createElement("li");
    const checkbox = document.createElement("input");
    checkbox.type = "checkbox";
    checkbox.checked = todo.Completed;
    const deleteBtn = document.createElement("button");
    deleteBtn.textContent = "Delete";
    
    checkbox.addEventListener("change", () => {
      fetch(`/todos/${todo.ID}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify( { Completed: checkbox.checked } )
      })
      .then(() => fetch("/todos"))
      .then(res => res.json())
      .then(renderTodos)
      .catch(err => console.error(err))
    });
    
    deleteBtn.addEventListener("click", () => {
      fetch(`/todos/${todo.ID}`, {
        method: "DELETE"
      })
      .then(() => fetch("/todos"))
      .then(res => res.json())
      .then(renderTodos)
      .catch(err => console.error(err));
    });
    
    li.appendChild(checkbox);
    li.appendChild(document.createTextNode(" " + todo.Description));
    li.appendChild(deleteBtn);
    list.appendChild(li);
  });
}
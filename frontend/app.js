window.addEventListener("load", () => {
    console.log("Page loaded")
    
    fetch("/todos")
    .then(response => {
        console.log("Raw response:", response);
        return response.json()
    })
    .then(data => {
        console.log("Todos from server:", data);
        renderTodos(data);
    })
    .catch(err => {
        console.error("Fetch error:", err);
    });
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
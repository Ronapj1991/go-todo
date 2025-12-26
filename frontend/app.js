window.addEventListener("load", () => {
    console.log("Page loaded")
    
    fetch("/todos")
    .then(response => {
        console.log("Raw response:", response);
        return response.json()
    })
    .then(data => {
        console.log("Todos from server:", data);
    })
    .catch(err => {
        console.error("Fetch error:", err);
    });
});
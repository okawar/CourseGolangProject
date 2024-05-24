document.getElementById("authForm").addEventListener("submit", function(event) {
    event.preventDefault();

    let formData = {
        email: null,
        login: document.getElementById("authLogin").value,
        fio: null,
        password: document.getElementById("authPassword").value,
        userid: null
    };

    fetch("api/auth", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.text())
    .then(data => {
        if(data.toString().includes("Success")) {
            window.location = "/index";
            localStorage.setItem("user_id", data.toString().split('=')[1]);
        } else {
            alert("Ошибка входа. Повторите попытку позже.");
        }
    })
    .catch(err => {
        console.error("Ошибка запроса авторизации:", err);
    })
});
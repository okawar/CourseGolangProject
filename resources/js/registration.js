document.getElementById("registrationForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Предотвращаем стандартное поведение отправки формы

    // Собираем данные из формы
    let formData = {
        email: document.getElementById("floatingInput").value,
        login: document.getElementById("floatingLogin").value,
        fio: document.getElementById("floatingFIO").value,
        password: document.getElementById("floatingPassword").value,
        confirmPassword: document.getElementById("confirmPassword").value
    };

    console.log("Отправляемые данные:", formData); // Добавляем лог для данных, отправляемых на сервер

    // Отправляем данные на сервер с использованием Fetch API
    fetch("/api/user/new", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => {
            console.log("Ответ сервера:", response);
            return response.json();
        }) // Преобразуем ответ в JSON
        .then(data => {
            console.log("Данные ответа:", data);
            if (data.success) {
                alert("Регистрация прошла успешно!");
                window.location.href = "/index";
            } else {
                alert("Произошла ошибка при регистрации: " + data.error);
            }
        })
        .catch(err => {
            console.error("Ошибка запроса на регистрацию:", err);
            alert("Произошла ошибка при регистрации. Пожалуйста, повторите попытку позже.");
        });
});

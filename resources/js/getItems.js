document.addEventListener("DOMContentLoaded", function() {
    //console.log(localStorage)
    fetch("/api/item/all")
        .then(response => {
            if (!response.ok) {
                switch (response.status) {
                    case 401:
                        window.location = "/start";
                        break;                
                    default:
                        alert("Ошибка запроса списка товаров. Перезагрузите страницу.");
                        break;
                }
            }
            return response.json();
        })
        .then(data => {
            //console.log("Data received from server:", data);
            const productsContainer = document.getElementById("productsContainer");

            data.forEach(product => {
                const card = createProductCard(product);
                productsContainer.appendChild(card);
            });
        })
        .catch(error => {
            console.error("Error:", error);
        });
});

function createProductCard(product) {
    const card = document.createElement("div");
    card.classList.add("col", "mt-5");
    card.innerHTML = `
    <div class="card shadow-sm">
      <img src="/resources/img/${product.item_id}.jpg">
      <div class="card-body">
        <h5 class="card-title">${product.name}</h5>
        <p class="card-text">${product.description}</p>
        <div class="d-flex justify-content-between align-items-center">
          <div class="btn-group">
            <button class="btn btn-sm btn-outline-secondary buy-button" onclick="addToCart(${product.item_id})">Купить</button>
          </div>
          <small class="text-body-secondary">${product.price} рублей</small>
        </div>
      </div>
    </div>
  `;
    return card;
}

function addToCart(itemId) {
    // Ensure the local storage has the required values
    const userId = Number(localStorage.getItem("user_id"));
    if (userId == null) {
        console.error("User ID is not available in localStorage");
        return;
    }

    fetch("/api/item/get?id=" + itemId.toString(), {
        method: "GET"
    })
        .then(response => response.json())
        .then(data => {
            localStorage.setItem("amount", data.price);
        });
        const price = Number(localStorage.getItem("amount"));
    // Prepare the data to be sent to the server
    let itemsData = {
        user_id: userId,
        item_id: itemId,
        amount: price
    };

    console.log("Prepared itemsData:", itemsData);

    // Convert to JSON string
    const jsonData = JSON.stringify(itemsData);
    console.log("JSON data to be sent:", jsonData);

    // Show confirmation modal
    const form = document.getElementById("addToCarts");
    const modal = new bootstrap.Modal(form);
    modal.show();
    form.addEventListener("submit", function(event) {
        event.preventDefault();
        fetch("/api/order/new", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: jsonData
        })
            .then(response => {
                console.log('Server response:', response);
                switch(response.status) {
                    case 200:
                        alert("Товар успешно добавлен в корзину.");
                        window.location = "/index";
                    break;
                    default:
                        alert("Ошибка отправки заказа. Попробуйте снова.");
                    break;
                }
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
}



document.addEventListener('DOMContentLoaded', function() {
  // Когда страница загружена, выполните запрос к серверу
  fetch("/api/order/all?owner_id=" + localStorage.getItem("user_id"))
      .then(response => response.text())
      .then(data => {
        for(order of JSON.parse(data)) {
          fetch("/api/item/get?id=" + order.item_id.toString())
          .then(response => response.json()) 
          .then(data => {
            const productsContainer = document.getElementById('productsContainer');
            const card = createProductCard(data);
            productsContainer.appendChild(card);
          });
        }
      });
});
// Функция для создания карточки товара
function createProductCard(product) {
  const card = document.createElement('div');
  card.classList.add('col', 'mt-5');
  card.innerHTML = `
    <div class="card shadow-sm">
      <img src="/resources/img/${product.item_id}.jpg">
      <div class="card-body">
        <h5 class="card-title">${product.name}</h5>
        <p class="card-text">${product.description}</p>
        <div class="d-flex justify-content-between align-items-center">
          <div class="btn-group">
            <!-- <button data-bs-toggle="modal" data-bs-target="#addToCarts" class="btn btn-sm btn-outline-secondary">Купить</button> -->
          </div>
          <small class="text-body-secondary">${product.price} рублей</small>
        </div>
      </div>
    </div>
  `;
  return card;
}

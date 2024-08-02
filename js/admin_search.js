function deleteBook() {
  const requestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ bookID: window.BookID }),
  };
  fetch("/admin/api/deletebooks", requestOptions).then((response) => {
    if (response.ok) {
      if (response.redirected) {
        window.location.href = response.url;
      }
    } else {
      response.json().then((error) => {
        const message = document.getElementById("message");
        message.innerText = error.message;
      });
    }
  });
}
const message = document.getElementById("message");
document.getElementById("updatebooks").addEventListener("submit", async (e) => {
  e.preventDefault();
  const title = document.getElementById("title").value;
  const genre = document.getElementById("genre").value;
  const author = document.getElementById("author").value;
  const dueTime = document.getElementById("duetime").value;
  const quantity = document.getElementById("quantity").value;
  const bookData = {
    bookID: window.BookID,
    title: title,
    genre: genre,
    author: author,
    dueTime: dueTime,
    quantity: quantity,
  };
  const requestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(bookData),
  };
  const response = await fetch("/admin/api/updatebooks", requestOptions);
  if (response.ok) {
    if (response.redirected) {
      window.location.href = response.url;
    } else {
      const error = await response.json();
      message.innerText = error.message;
    }
  } else {
    const error = await response.json();
    message.innerText = error.message;
  }
});

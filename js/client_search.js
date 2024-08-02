const message = document.getElementById("message");

function CheckOut() {
  const bookID = BookID;
  const RequestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ bookID }),
  };
  fetch("/client/api/checkout", RequestOptions).then(async (response) => {
    if (response.ok) {
      const res = await response.json().then((data) => data.message);
      console.log(res);
      alert(res);
      window.location.reload();
    } else {
      await response.json().then((error) => {
        message.innerText = error.message;
      });
    }
  });
}

function CheckIn() {
  const bookID = BookID;
  const RequestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ bookID }),
  };
  fetch("/client/api/checkin", RequestOptions).then(async (response) => {
    if (response.ok) {
      const res = await response.json().then((data) => data.message);
      alert(res);
      window.location.reload();
    } else {
      await response.json().then((error) => {
        message.innerText = error.message;
      });
    }
  });
}

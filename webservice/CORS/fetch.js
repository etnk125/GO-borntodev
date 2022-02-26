//ngrok public api
const URL =
  "https://53ab-2001-44c8-43a4-93d4-99d-a014-4135-f711.ngrok.io/user/1";
fetch(URL)
  .then((response) => response.json())
  .then((data) => {
    console.log(data);
  });

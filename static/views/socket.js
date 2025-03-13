export const socket = new WebSocket("ws://localhost:8080/ws");



// //session scroll
// export async function renderChatEvents() {
//     const message_Box = document.querySelector(".ChatBody");
//     message_Box.addEventListener("scroll", function (e) {
//       if (e.target.scrollTop === 0) {
//         let scroll = parseInt(sessionStorage.getItem("scroll")) + 10;
//         sessionStorage.setItem("scroll", scroll);
//         throttle(LoadMessages(), 3000);
//       }
//     });
// }
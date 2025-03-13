export function sessions(data) {
  //sessionStorage.setItem("UserId", data.UserId);
  sessionStorage.setItem("UserId", data.UserId);
  sessionStorage.setItem("Username", data.Username);
  sessionStorage.setItem("ProfileImg", data.ProfileImg);
  const usersData = data.Users;
  // Convert the array to a JSON string
  const usersDataJSON = JSON.stringify(usersData);
  // Store the users data in session storage
  sessionStorage.setItem("users", usersDataJSON);
  sessionStorage.setItem("SessionId", data.SessionId);
}

export function setSessionCookie(userSessionUuid) {
  removeSessionCookie();
  const expires = new Date(Date.now() + 3600 * 1000); // Expires in 1 hour
  const cookieString = `session_id=${userSessionUuid}; expires=${expires.toUTCString()}; domain=localhost; path=/; SameSite=None; Secure=true`;
  console.log(cookieString);
  document.cookie = cookieString;
}

// check if there is a session cookie
export function getSessionCookie() {
  const cookieValue = document.cookie
    .split("; ")
    .find((row) => row.startsWith("session_id"))
    .split("=")[1];
  return cookieValue;
}

// remove the session cookie
export function removeSessionCookie() {
  document.cookie =
    "session_id=; expires=Thu, 01 Jan 1970 00:00:00 UTC; domain=localhost; path=/; SameSite=None; Secure=true";
}

function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(";").shift();
}
const cookie = getCookie("session_id");


export function testCookie(){
  const allCookies = document.cookie;
  console.log(allCookies)
  // Split the cookies into an array
  const cookiesArray = allCookies.split(';');
  
  // Loop through the array to find the specific cookie
  let sessionCookieValue = null;
  cookiesArray.forEach(cookie => {
      const [name, value] = cookie.trim().split('=');
      if (name === 'userID') {
          sessionCookieValue = value;
      }
  });
  
  console.log('Session Cookie Value:', sessionCookieValue);
  return sessionCookieValue
}
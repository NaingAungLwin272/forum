importScripts(
  "https://www.gstatic.com/firebasejs/10.3.0/firebase-app-compat.js"
);
importScripts(
  "https://www.gstatic.com/firebasejs/10.3.0/firebase-messaging-compat.js"
);

firebase.initializeApp({
  apiKey: "AIzaSyC8anpH9gnHBlejfKV7EasUWTAaQGCYi9c",
  authDomain: "mtmcommunity-cd677.firebaseapp.com",
  projectId: "mtmcommunity-cd677",
  storageBucket: "mtmcommunity-cd677.appspot.com",
  messagingSenderId: "9291544284",
  appId: "1:9291544284:web:cab6743ffd7bd99a9b2619",
  measurementId: "G-0L25G3Y0TP",
});
const messaging = firebase.messaging();

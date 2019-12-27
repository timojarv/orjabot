import * as firebase from 'firebase';

import 'firebase/auth';
import 'firebase/firestore';

const firebaseConfig = {
    apiKey: "AIzaSyB8GIyNQeGVqNZrbFHPMo-zopOy72TuWUw",
    authDomain: "orjabot.firebaseapp.com",
    databaseURL: "https://orjabot.firebaseio.com",
    projectId: "orjabot",
    storageBucket: "orjabot.appspot.com",
    messagingSenderId: "170626747084",
    appId: "1:170626747084:web:d5f423d3382fb404741069"
};

firebase.initializeApp(firebaseConfig);


const provider = new firebase.auth.GoogleAuthProvider();
firebase.auth().useDeviceLanguage();

export const login = () => firebase.auth().signInWithPopup(provider);
export const logout = () => firebase.auth().signOut();

export const db = firebase.firestore();

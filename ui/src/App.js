import './App.css';

import React, { useState, useEffect } from 'react';
import axios from 'axios';

function App() {
  const [movies, setMovies] = useState([]);

  const [title, setTitle] = useState("");
  const [director, setDirector] = useState("");
  const [message, setMessage] = useState("");

  let handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let res = await fetch(process.env.REACT_APP_BACKEND_URL+"/movies", {
        method: "POST",
        body: JSON.stringify({
          title: title,
          director: director
        }),
        mode: 'no-cors',
      });
      let resJson = await res.json();

      if (res.status === 201) {
        setTitle("");
        setDirector("");
        setMessage("User created successfully");
      } else {
        setMessage("Some error occured");
      }

    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    axios.get(process.env.REACT_APP_BACKEND_URL+"/movies", {mode: 'no-cors'})
      .then(response => {
        setMovies(response.data);
        console.log("123", response)
      })
      .catch(error => {
        console.error(error);
      });
  }, []);

  return (
    <div style={{
      position: 'relative',
      margin: 'auto',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    }}>
      <div>
        <h1>Awesome Movie App</h1>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={title}
            placeholder="Title"
            onChange={(e) => setTitle(e.target.value)}
          />
          <input
            type="text"
            value={director}
            placeholder="Director"
            onChange={(e) => setDirector(e.target.value)}
          />
      
          <button type="submit">Add</button>

          <div className="message">{message ? <p>{message}</p> : null}</div>
        </form>
        <ul>
          {movies.map(entry => (
            <li key={entry.ID}>{entry.title} - {entry.director}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;

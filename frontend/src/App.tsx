import { useEffect, useState } from 'react'
import { Link } from 'react-router'
import './App.css'

function App() {
  const [hotels, setHotels] = useState([])

  useEffect(() => {
    const fetchData = () => {
      fetch('http://localhost:8080/hotels')
      .then(response => response.json())
      .then(data => setHotels(data))
      .catch(error => console.error('Error fetching data:', error))
    }
    fetchData();
  }, []);

  return (
    <>
      <h1>Hotels</h1>
        
      {hotels.map(hotel => (
        <div>
          <h2><Link to={`hotels/${hotel.id}`}>{hotel.name}</Link></h2>
          <p>{hotel.location}</p>
        </div>
      ))}
    </>
  )
}

export default App

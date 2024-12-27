import { useEffect, useState } from 'react'
import { Link } from 'react-router'

function HotelsPage() {
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
      <h3>Hotels</h3>

      {hotels.map(hotel => (
        <article key={hotel.id}>
          <h5><Link to={`${hotel.id}`}>{hotel.name}</Link></h5>
          <div>{hotel.location}</div>
        </article>
      ))}
    </>
  )
}

export default HotelsPage

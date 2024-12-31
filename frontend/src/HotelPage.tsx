import { useEffect, useState } from 'react'
import { useParams } from "react-router";

function HotelPage() {
  const { hotelId } = useParams();
  const [hotel, setHotel] = useState({})
  const [rooms, setRooms] = useState([])

  useEffect(() => {
    const fetchHotel = () => {
      fetch(`http://localhost:8080/hotels/${hotelId}`)
        .then(response => response.json())
        .then(data => setHotel(data))
        .catch(error => console.error('Error fetching data:', error))
    }
    const fetchHotelRooms = () => {
        fetch(`http://localhost:8080/hotels/${hotelId}/rooms`)
          .then(response => response.json())
          .then(data => setRooms(data))
          .catch(error => console.error('Error fetching data:', error))
      }
    fetchHotel();
    fetchHotelRooms();
  }, [hotelId]);

  return (
    <>
      <h3>{hotel.name}</h3>
        
      {rooms.map(room => (
        <article key={room.id}>
          <h5>from the movie "{room.title}"</h5>
          <div>{room.type}</div>
          <div>{room.description}</div>
          <div>Size: {room.size} m2</div>
          <div>Price per night: <strong>{room.price}€</strong></div>
        </article>
      ))}
    </>
  )
}

export default HotelPage

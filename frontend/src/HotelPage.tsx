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
      <h1>{hotel.name}</h1>
        
      {rooms.map(room => (
        <div>
          <h2>from the movie "{room.title}"</h2>
          <p>{room.type}</p>
          <p>Size: {room.size} m2</p>
          <p>{room.description}</p>
        </div>
      ))}
    </>
  )
}

export default HotelPage

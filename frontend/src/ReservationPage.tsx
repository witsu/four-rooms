import { useState } from 'react'
import { useSearchParams } from "react-router";

function SearchPage() {
  const [reservation, setReservation] = useState();
  const [searchParams] = useSearchParams();
  const start = searchParams.get("start");
  const end = searchParams.get("end");
  const roomId = searchParams.get("roomId");

  function reserve(event) {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const data = {
      first_name: formData.get("first_name"),
      last_name: formData.get("last_name"),
      email: formData.get("email"),
      room_id: Number(roomId),
      start_date: start,
      end_date: end,
    };
    fetch('http://localhost:8080/reservations', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => setReservation(data))
    .catch(error => console.error('Error sending data:', error))

  }

  return (
    <>
      {!reservation && (
      <div>
        <h3>Reservation</h3>
        <p>Fill in the form to book a room {roomId} from {start} to {end}.</p>
      
        <form onSubmit={reserve}>
          <div className="field label border">
            <input type="text" name="first_name" />
            <label>First name</label>
          </div>
          <div className="field label border">
            <input type="text" name="last_name" />
            <label>Last name</label>
          </div>
          <div className="field label border">
            <input type="text" name="email" />
            <label>Email</label>
          </div>
          <button type="submit">Book</button>
        </form>
      </div>
      )}
      {reservation && (
      <div>
        <h3>Reservation confirmed</h3>
        <p>Your reservation number {reservation.id} has been confirmed. You will receive an email with the details.</p>
      </div>
      )}
    </>
  )
}

export default SearchPage

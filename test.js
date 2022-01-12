const data = {
  photographers: [
    {
      id: "1",
      name: "Otto Crawford",
      availabilities: [
        {
          starts: "2020-11-25T08:00:00.000Z",
          ends: "2020-11-25T16:00:00.000Z",
        },
      ],
      bookings: [
        {
          id: "1",
          starts: "2020-11-25T08:30:00.000Z",
          ends: "2020-11-25T09:30:00.000Z",
        },
      ],
    },
    {
      id: "2",
      name: "Jens Mills",
      availabilities: [
        {
          starts: "2020-11-25T08:00:00.000Z",
          ends: "2020-11-25T09:00:00.000Z",
        },
        {
          starts: "2020-11-25T13:00:00.000Z",
          ends: "2020-11-25T16:00:00.000Z",
        },
      ],
      bookings: [
        {
          id: "2",
          starts: "2020-11-25T15:00:00.000Z",
          ends: "2020-11-25T16:00:00.000Z",
        },
      ],
    },
  ],
};

const res = [
  {
    photographer: {
      id: 1,
      name: "Otto Crawford",
    },
    timeSlot: {
      starts: "2020-11-25T09:30:00.000Z",
      ends: "2020-11-25T11:00:00.000Z",
    },
  },
  {
    photographer: { id: "2", name: "Jens Mills" },
    timeSlot: {
      starts: "2020-11-25T13:00:00.000Z",
      ends: "2020-11-25T14:30:00.000Z",
    },
  },
];

function availableTimeSlotsForBooking(durationInMinutes) {
  durationInMinutes = Number(durationInMinutes);
  const availablePhotographers = [];
  for (const photographer of data.photographers) {
    const timeSlot = photographer.availabilities.find((availability) => {
      const availabilityStarts = new Date(availability.starts).getTime();
      const availabilityEnds = new Date(availability.ends).getTime();
      const timeDiff =
        new Date(availability.ends).getTime() -
        new Date(availability.starts).getTime();

      const totalAvailablityMinutes = Math.ceil(timeDiff / (1000 * 60));
      const bookingsInSlot = photographer.bookings.filter((booking) => {
        return (
          new Date(booking.starts).getTime() >= availabilityStarts &&
          availabilityEnds >= new Date(booking.ends).getTime()
        );
      });

      let availabilityHasEnoughTime = false;
      if (totalAvailablityMinutes >= durationInMinutes) {
        if (bookingsInSlot.length) {
          // There is a booking for this available time
          availabilityHasEnoughTime = bookingsInSlot.find((booking) => {
            // Get the total minutes of the time
            const bookingTotalTime =
              new Date(booking.ends).getTime() -
              new Date(booking.starts).getTime();
            const bookingTotalMin = Math.ceil(bookingTotalTime / (1000 * 60));
            // Remove the booking time from the availibity time
            // Then check if the passed duration can fit in to the remaining time after the time for the current booking has been removed.
            // if there is enough space left, the user can share with the current booking.
            return (
              totalAvailablityMinutes - bookingTotalMin >= durationInMinutes
            );
          });
        }
      }
      return availabilityHasEnoughTime ? availability : undefined;
    });

    if (timeSlot) {
      const { id, name } = photographer;
      availablePhotographers.push({ timeSlot, photographer: { id, name } });
    }
  }
  return availablePhotographers;
}

console.log(JSON.stringify(availableTimeSlotsForBooking("19")));

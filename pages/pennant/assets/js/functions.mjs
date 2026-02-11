// Calculate the last day of the week
function getLastDayOfWeek(date) {
  const lastDay = new Date(date);
  // getDay() returns 0 for Sunday, 1 for Monday, etc.
  // Sunday is typically the last day of the week (default in flatpickr)
  const lastDayIndex = 6;
  const currentDayIndex = lastDay.getDay();
  const daysToAdd = lastDayIndex - currentDayIndex;

  lastDay.setDate(lastDay.getDate() + daysToAdd);
  return lastDay;
}

// Initializes week-ending flatpickr
const initPennantPage = (event) => {
  var fp = flatpickr("#week-ending", {
    appendTo: document.getElementById("contains-flatpickr"),
    dateFormat: "j M Y",
    enableTime: false,
    plugins: [new weekSelect({})],
    onChange: function (selectedDates, _dateStr, instance) {
      if (selectedDates.length > 0) {
        let weekEnd = getLastDayOfWeek(selectedDates[0]);
        instance.input.value = instance.formatDate(weekEnd, "j M Y");
      }
    },
  });

  // Initialize the date picker with the current date and give it input focus
  let today = new Date();
  fp.setDate(getLastDayOfWeek(today));

  let ct = document.getElementById("competition-selector");
  if (ct) {
    ct.focus();
  }
};

export { initPennantPage };

window.initPennantPage = initPennantPage; // Expose the function to the global scope

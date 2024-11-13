export const dateStrtoDDMMYYYY = (dateStr) => {
  const date = new Date(dateStr);

  // Check if the date is valid
  if (isNaN(date.getTime())) {
    throw new Error("Invalid date format");
  }

  const formatted = `${String(date.getDate()).padStart(2, "0")}-${String(
    date.getMonth() + 1
  ).padStart(2, "0")}-${date.getFullYear()}`;
  return formatted;
};

export const ddMMYYYtoMMDDYYYY = (dateStr) => {
  console.log(dateStr);
  const [day, month, year] = dateStr.split("-");
  return `${month}-${day}-${year}`;
};

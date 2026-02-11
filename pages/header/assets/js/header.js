function capitalizeFirstLetter(string) {
  if (typeof string !== "string" || string.length === 0) {
    return string; // Handle empty strings or non-string inputs
  }
  return string.charAt(0).toUpperCase() + string.slice(1);
}

function initPage(name, event) {
  fmtName = "init" + capitalizeFirstLetter(name.replace(/ /g, "")) + "Page";

  if (typeof window[fmtName] === "function") {
    console.log("Executing initialization function: ", fmtName);
    window[fmtName](event); // Execute the initialization function
  } else {
    console.log("Initialization function not found: ", fmtName);
  }
}

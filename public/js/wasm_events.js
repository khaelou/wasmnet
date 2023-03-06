// WASM Events
var textarea;

async function addPoolOutput(targetLine) {
    textarea = document.getElementById("poolOutput");
    textarea.value += targetLine + "\n";
    textarea.scrollTop = textarea.scrollHeight;
}

async function addForInput(targetLine) {
    textarea = document.getElementById("forInput");
    textarea.value += targetLine + "\n";
    textarea.scrollTop = textarea.scrollHeight;
}

async function addForConsoleOutput(targetLine) {
    textarea = document.getElementById("forOutput");
    textarea.value += targetLine;
    textarea.scrollTop = textarea.scrollHeight;
}
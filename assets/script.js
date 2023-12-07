const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
.then((result) => {
    goWasm.run(result.instance)
    // document.getElementById("submitLogin").addEventListener("click", function(event) {
    //     event.preventDefault();
    
    //     const login = document.getElementById("loginUsername").value;
    //     const password = document.getElementById("loginPassword").value;
    
    //     const data = {
    //         login: login,
    //         password: password
    //     };
    
    //     fetch("http://localhost:8083/login", {
    //         method: "POST",
    //         headers: {
    //             "Content-Type": "application/json"
    //         },
    //         body: JSON.stringify(data)
    //     })
    //     .then(response => response.json())
    //     .then(data => {
    //         console.log("Login successful:", data);
    //         // Handle successful login response
    //     })
    //     .catch(error => {
    //         console.error("Error during login:", error);
    //         // Handle login error
    //     });
    // });    
})


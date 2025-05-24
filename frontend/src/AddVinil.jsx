import { useState } from "react";

const AddVinil = () => {

    const [title, setTitle] = useState("");
    const [artist, setArtist] = useState("");
    const [year, setYear] = useState("");
    const [price, setPrice] = useState("");
    const [error, setError] = useState("");

    const handleSave = async (e) => {
        e.preventDefault();
        setError("");
    

    try {    
            const response = await fetch("http://localhost:8080/albums", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ title, artist, year: parseInt(year), price: parseFloat(price)})
            });   //necesito esperar el post al servicio
    
            if (!response.ok) throw new Error("Error al guardar un nuevo album");
    } catch (error) {
        setError("Error al guardar nuevo album")
    
    }
}

    return (
        <div>
            <form onSubmit={handleSave}>

                <input 
                    placeholder="Titulo del album"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />

                <input 
                    placeholder="Artista"
                    value={artist}
                    onChange={(e) => setArtist(e.target.value)}
                    required
                />

                <input 
                    placeholder="Anio"
                    value={year}
                    onChange={(e) => setYear(e.target.value)}
                    required
                />

                <input 
                    placeholder="Precio"
                    value={price}
                    onChange={(e) => setPrice(e.target.value)}
                    required
                />
                <button type="submit">Guardar</button>
            </form>
        </div>
    )
}

export default AddVinil;

function add_new_img_field() {
    let gallery = document.getElementById("gallery")
    let gal = gallery.children.length/2

    let src_div = document.createElement("div")
    src_div.className = "row"

    let src_label_div = document.createElement("div")
    src_label_div.className = "col-25"

    let src_input_div = document.createElement("div")
    src_input_div.className = "col-75"

    let src_label = document.createElement("label")
    src_label.setAttribute("for", "src"+ gal.toString())
    src_label.innerHTML = "Source of image "+gal.toString()

    let src_input = document.createElement("input")
    src_input.setAttribute("type", "text")
    src_input.id = "src"+gal.toString()
    src_input.name = "src"+gal.toString()
    src_input.value = "/docimg/Attractions/"
    src_input.required = true

    src_label_div.appendChild(src_label)
    src_input_div.appendChild(src_input)
    src_div.appendChild(src_label_div)
    src_div.appendChild(src_input_div)
    document.getElementById("gallery").appendChild(src_div)

    let cap_div = document.createElement("div")
    cap_div.className = "row"

    let cap_label_div = document.createElement("div")
    cap_label_div.className = "col-25"

    let cap_input_div = document.createElement("div")
    cap_input_div.className = "col-75"

    let cap_label = document.createElement("label")
    cap_label.setAttribute("for", "cap"+ gal.toString())
    cap_label.innerHTML = "Caption of image "+gal.toString()

    let cap_input = document.createElement("input")
    cap_input.setAttribute("type", "text")
    cap_input.id = "cap"+gal.toString()
    cap_input.name = "cap"+gal.toString()
    cap_input.required = true

    cap_label_div.appendChild(cap_label)
    cap_input_div.appendChild(cap_input)
    cap_div.appendChild(cap_label_div)
    cap_div.appendChild(cap_input_div)
    document.getElementById("gallery").appendChild(cap_div)

    gal++
}
function remove_last_img_field(){
    let gallery = document.getElementById("gallery")
    if(gallery.children.length){
        for(let i = 0; i < 2; i++) {
            let last_img_field = gallery.lastChild
            gallery.removeChild(last_img_field)
        }
    }
}
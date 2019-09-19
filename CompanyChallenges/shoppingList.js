function shoppingList(items) {
matches = items.match(/[0-9.]+/g)
    console.log(matches)
    total = 0
    if(matches){
        for(i=0; i < matches.length;i++){
            matches[i] = matches[i].replace(/\.\./, "")
            console.log(matches[i])
            total = parseFloat(total) + parseFloat(matches[i])
        }
    }
    return total
}

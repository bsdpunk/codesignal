function fancyRide(l, fares) {
max = 0
    index =0
    opt = [
        {"option": "UberX"},
        {"option": "UberXL"},
        {"option": "UberPlus"},
        {"option": "UberBlack"},
        {"option": "UberSUV"}
    ]
    
    for(i=0; i < fares.length; i++){
    if( (l * fares[i]) <= 20 && max <  (l * fares[i]) ){
        max =  (l * fares[i])
        index = i
    }
}
    return opt[index]["option"]
}

function incorrectPasscodeAttempts(passcode, attempts) {
tries = 0
    for(i=0;i < attempts.length;i++){
    if(attempts[i] == passcode){
        tries =0
        
    }else{
        tries = tries + 1
        if(tries > 9){
            return true
        }
    }
}
    return false
}

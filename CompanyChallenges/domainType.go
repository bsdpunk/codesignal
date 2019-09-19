import "regexp"
func domainType(domains []string) []string {

    tlds := map[string]string{
    "info": "information",
    "com" : "commercial",
    "net" : "network",
    "org" : "organization",
    
}
    var final []string
    re := regexp.MustCompile(`[a-z]{3,4}$`)
    for i := range domains {
        fmt.Println(re.FindStringSubmatch(domains[i]))
        final = append(final, tlds[string(re.FindStringSubmatch(domains[i])[0])])
    } 
    return final
}



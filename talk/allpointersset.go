// START OMIT
func main() {
    port := 20
    str := "hello"
    config := {
        Port: &port,
        str: &str,
    }

    strawman,_ := NewServer(addr, config)
// END OMIT
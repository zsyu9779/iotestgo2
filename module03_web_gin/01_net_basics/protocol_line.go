package main

import (
    "bufio"
    "errors"
    "net"
)

func readLine(conn net.Conn, max int) (string, error) {
    r := bufio.NewReader(conn)
    s, err := r.ReadString('\n')
    if err != nil {
        return "", err
    }
    if len(s) > max {
        return "", errors.New("line too long")
    }
    return s[:len(s)-1], nil
}

func writeLine(conn net.Conn, s string) error {
    _, err := conn.Write([]byte(s + "\n"))
    return err
}

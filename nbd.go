package qemu

import (
	"fmt"
	"os"
	"os/exec"
)

// ConnectNbd FILE to the local NBD device DEV
// qemu-nbd --connect=/dev/nbd0 image.qcow2
func ConnectNbd(dev string, file string) (error) {
    if _, err := os.Stat("/var/lock/qemu-nbd-" + dev) ; os.IsNotExist(err) {
        cmd := exec.Command("sudo", "qemu-nbd", "--connect=/dev/" + dev, file)
        out, err := cmd.CombinedOutput()
        if err != nil {
		    return fmt.Errorf("'qemu-nbd --connect' output: %s", oneLine(out))
	    }
        return nil
    }
    return fmt.Errorf("Cannot connect on locked device /var/lock/qemu-nbd-"+ dev)
}

// DisconnectNbd: qemu-nbd --disconnect /dev/nbd0 
func DisconnectNbd(dev string) (error) {
    if _, err := os.Stat("/var/lock/qemu-nbd-"+ dev) ; os.IsNotExist(err) {
        return fmt.Errorf("Cannot disconnect not locked device /var/lock/qemu-nbd-"+ dev)
    }
    cmd := exec.Command("sudo", "qemu-nbd", "--disconnect", "/dev/" + dev)
    out, err := cmd.CombinedOutput()
    if err != nil {
		return fmt.Errorf("'qemu-nbd --disconnect' output: %s", oneLine(out))
	}
    return nil
}

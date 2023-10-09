package types

import "errors"

type Runtime string

const (
	Docker     Runtime = "docker"
	KVM        Runtime = "kvm"
	Xen        Runtime = "xen"
	Containerd Runtime = "containerd"
	Kubernetes Runtime = "kubernetes"
)

func (runtime Runtime) ToString() string {
	switch runtime {
	case Docker:
		return "docker"
	case KVM:
		return "kvm"
	case Containerd:
		return "containerd"
	case Kubernetes:
		return "kubernetes"
	case Xen:
		return "xen"
	default:
		return ""
	}
}

func ToRuntime(input string) (Runtime, error) {
	switch input {
	case "docker":
		return Docker, nil
	case "kvm":
		return KVM, nil
	case "containerd":
		return Containerd, nil
	case "kubernetes":
		return Kubernetes, nil
	case "xen":
		return Xen, nil
	default:
		return "", errors.New("Unsupported runtime: " + input)
	}
}

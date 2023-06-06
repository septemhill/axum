package axum

import "net/http"

type Service interface {
	Handle(*http.Request) ResponsePacker
}

type ServiceFunc func(*http.Request) ResponsePacker

func (h ServiceFunc) Handle(r *http.Request) ResponsePacker {
	return h(r)
}

type Arg1[T1 RequestExtractor] func(*T1) ResponsePacker

func (a Arg1[T1]) Handle(req *http.Request) ResponsePacker {
	var ext T1
	t, err := ext.Extract(req)
	if err != nil {
		// TODO: return extract failed error response packer
		return InvalidParameter(100, err.Error())
	}

	s := (t.(T1))
	return a(&s)
}

type Arg2[T1, T2 RequestExtractor] func(*T1, *T2) ResponsePacker

func (a Arg2[T1, T2]) Handle(req *http.Request) ResponsePacker {
	var ext1 T1
	t1, err := ext1.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s1 := (t1.(T1))

	var ext2 T2
	t2, err := ext2.Extract(req)
	if err != nil {
		// TODO
		return InvalidParameter(100, err.Error())
	}
	s2 := (t2.(T2))

	return a(&s1, &s2)
}

type Arg3[T1, T2, T3 RequestExtractor] func(*T1, *T2, *T3) ResponsePacker

func (a Arg3[T1, T2, T3]) Handle(req *http.Request) ResponsePacker {
	var ext1 T1
	t1, err := ext1.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s1 := (t1.(T1))

	var ext2 T2
	t2, err := ext2.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s2 := (t2.(T2))

	var ext3 T3
	t3, err := ext3.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s3 := (t3.(T3))

	return a(&s1, &s2, &s3)
}

type Arg4[T1, T2, T3, T4 RequestExtractor] func(*T1, *T2, *T3, *T4) ResponsePacker

func (a Arg4[T1, T2, T3, T4]) Handle(req *http.Request) ResponsePacker {
	var ext1 T1
	t1, err := ext1.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s1 := (t1.(T1))

	var ext2 T2
	t2, err := ext2.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s2 := (t2.(T2))

	var ext3 T3
	t3, err := ext3.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s3 := (t3.(T3))

	var ext4 T4
	t4, err := ext4.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s4 := (t4.(T4))

	return a(&s1, &s2, &s3, &s4)
}

type Arg5[T1, T2, T3, T4, T5 RequestExtractor] func(*T1, *T2, *T3, *T4, *T5) ResponsePacker

func (a Arg5[T1, T2, T3, T4, T5]) Handle(req *http.Request) ResponsePacker {
	var ext1 T1
	t1, err := ext1.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s1 := (t1.(T1))

	var ext2 T2
	t2, err := ext2.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s2 := (t2.(T2))

	var ext3 T3
	t3, err := ext3.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s3 := (t3.(T3))

	var ext4 T4
	t4, err := ext4.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s4 := (t4.(T4))

	var ext5 T5
	t5, err := ext5.Extract(req)
	if err != nil {
		return InvalidParameter(100, err.Error())
	}
	s5 := (t5.(T5))

	return a(&s1, &s2, &s3, &s4, &s5)
}

func Arg1Func[
	T1 RequestExtractor,
](f Arg1[T1]) Arg1[T1] {
	return f
}

func Arg2Func[T1, T2 RequestExtractor,
](f Arg2[T1, T2]) Arg2[T1, T2] {
	return f
}

func Arg3Func[T1, T2, T3 RequestExtractor,
](f Arg3[T1, T2, T3]) Arg3[T1, T2, T3] {
	return f
}

func Arg4Func[T1, T2, T3, T4 RequestExtractor,
](f Arg4[T1, T2, T3, T4]) Arg4[T1, T2, T3, T4] {
	return f
}

func Arg5Func[T1, T2, T3, T4, T5 RequestExtractor,
](f Arg5[T1, T2, T3, T4, T5]) Arg5[T1, T2, T3, T4, T5] {
	return f
}

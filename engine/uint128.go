package engine

type uint128 struct {
	hi uint64
	lo uint64
}

func mask6(n int) uint128 {
	return uint128{^(^uint64(0) >> n), ^uint64(0) << (128 - n)}
}

func (u uint128) isZero() bool { return u.hi|u.lo == 0 }

func (u uint128) and(m uint128) uint128 {
	return uint128{u.hi & m.hi, u.lo & m.lo}
}

func (u uint128) xor(m uint128) uint128 {
	return uint128{u.hi ^ m.hi, u.lo ^ m.lo}
}

func (u uint128) or(m uint128) uint128 {
	return uint128{u.hi | m.hi, u.lo | m.lo}
}

func (u uint128) not() uint128 {
	return uint128{^u.hi, ^u.lo}
}

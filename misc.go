package errgo

type addDetails func(k, v string)

func setDetails(k, v string, d *Details) {
	if v != "" {
		d.Add(k, v)
	}
}

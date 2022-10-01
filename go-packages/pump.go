package gopackages

type Pump struct{
	state bool
}

func (p *Pump) OpenPump(){
	p.state = true;
}

func (p *Pump) ClosePump(){
	p.state = false;
}


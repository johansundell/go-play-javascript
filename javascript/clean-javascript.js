function term(n)
{
	var i = (n) * 2 + 1;
	if(n % 2  ===  0) 
		return -4 / ((i - 1) * i * (i + 1));
	else
		return 4 / ((i - 1) * i * (i + 1));
}

function GetPi(n)
{
	//var f, i, n;
	var f = 3.0;
	var i = 1;
	while (true) {
		if (!(i <= n))
			break;
		f = f + (term(i));
		i++;
	}
	return f;
}
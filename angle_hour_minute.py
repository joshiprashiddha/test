def calcAngle(h, m):
    if (m < 0 or h < 0 or not(isinstance(h,int) & isinstance(m,int))):
        print("Can process only positive integer values for hour and minute")
        return

    # for any positive minute value greater or equal to 60
    if m >= 60:
        m = m//60
        h = h + (m%60)

    # for any positive hour value greater or equal to 60
    if h >= 12:
        h = h//12

    # total hour travelled.
    total_hour = h + m/60
    h_degree_with_respect_to_12 = (360/12)*total_hour

    # total minute travelled.
    m_degree_with_respect_to_12 = (360/60)*m

    total_degree_diff = abs(h_degree_with_respect_to_12-m_degree_with_respect_to_12)
    if(total_degree_diff >=180 ):
        total_degree_diff = 360 - total_degree_diff
    return total_degree_diff

def main():
    print(calcAngle(11,30))

main()
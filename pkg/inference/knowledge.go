package inference

// Simple internationalization, get TB description.
func getTBDescription(locale string) string {
	if locale == "id" {
		return "Tuberculosis (TB) adalah infeksi bakteri yang menyebar dengan cara menghirup droplet ringan dari bekas batuk atau bersin dari seseorang yang sudah terinfeksi. Biasanya, TB menyerang paru-paru, tetapi TB bisa menyerang bagian tubuh lain, seperti perut, kelenjar, tulang, dan sistem saraf pusat. TB adalah kondisi yang berpotensi menjadi serius, tetapi dapat diobati apabila menggunakan antibiotik yang benar. TB menyebar melalui udara. Ketika orang yang menderita TB paru-paru batuk, meludah, atau bersin, mereka mengeluarkan bakteri TB ke udara. Seseorang hanya perlu menghirup sedikit dari bakteri ini untuk menjadi terinfeksi. Sekitar satu per empat populasi dunia menderita TB, yang menandakan bahwa ada orang sudah terinfeksi dengan bakteri TB tetapi belum sakit dan tidak bisa menularkannya."
	}

	return "Tuberculosis (TB) is a bacterial infection spread through inhaling tiny droplets from the coughs or sneezes of an infected person. It mainly affects the lungs, but it can affect any part of the body, including the tummy (abdomen), glands, bones and nervous system. TB is a potentially serious condition, but it can be cured if it is treated with the right antibiotics. TB is spread from person to person through the air. When people with lung TB cough, sneeze or spit, they propel the TB germs into the air. A person needs to inhale only a few of these germs to become infected. About one-quarter of the world's population has a TB infection, which means people have been infected by TB bacteria but are not (yet) ill with the disease and cannot transmit it."
}

// Simple internationalization, get TB treatment.
func getTBTreatment(locale string) string {
	if locale == "id" {
		return "TB adalah penyakit yang bisa disembuhkan. Biasanya, diperlukan antibiotik yang berdurasi enam bulan. Biasanya, ada empat buah antibiotik yang diberikan dengan informasi dan dukungan pada pasien oleh petugas kesehatan. Tanpa antibiotik, pengobatan TB menjadi lebih sulit. Sejak 2000, estimasi 66 juta jiwa sudah terselamatkan dari TB karena melakukan diagnosis dan perawatan."
	}

	return "TB is a treatable and curable disease. Active, drug-susceptible TB disease is treated with a standard 6-month course of 4 antimicrobial drugs that are provided with information and support to the patient by a health worker or trained volunteer. Without such support, treatment adherence is more difficult. Since 2000, an estimated 66 million lives were saved through TB diagnosis and treatment."
}

// Simple internationalization, get TB prevention.
func getTBPrevention(locale string) string {
	if locale == "id" {
		return "Anda dapat melakukan pencegahan dengan cara menyiapkan ventilasi yang baik, cahaya natural, dan menjaga agar kondisi lingkungan Anda tetap bersih. Direkomendasikan untuk mendapatkan vaksin untuk TB sebagai salah satu pencegahan yang paling efektif. Vaksin membuat sistem imun Anda lebih kuat, yang dapat membuat Anda menjadi lebih terjaga dari terserang TB (ada yang hingga 15 tahun). Karena beberapa faktor eksternal, ada beberapa orang yang memiliki risiko yang lebih besar untuk terkena dan sakit karena TB."
	}

	return "You can perform several precautions by providing good ventilation, natural light, and keeping everything clean. It is recommended to take vaccinations of TB as well. Vaccinations help you to keep your immune system in prime condition, thus allowing you to resist the virus for a longer time (some up to 15 years). Depending on some external factors, some people are more at risk for being exposed into developing the TB disease."
}

// All diseases that are in this expert system.
// Data is processed from https://www.kaggle.com/victorcaelina/tuberculosis-symptoms.
func getDiseases(locale string) []Disease {
	diseases := []Disease{
		{
			ID:          "D01",
			Name:        "Tuberculosis",
			Description: getTBDescription(locale),
			Treatment:   getTBTreatment(locale),
			Prevention:  getTBPrevention(locale),
			Source: []SourceAndLink{
				{
					Name: "NHS",
					Link: "https://www.nhs.uk/conditions/tuberculosis-tb/",
				},
				{
					Name: "WHO",
					Link: "https://www.who.int/news-room/fact-sheets/detail/tuberculosis",
				},
				{
					Name: "TBAlert",
					Link: "https://www.tbalert.org/about-tb/what-is-tb/prevention/",
				},
				{
					Name: "CDC Government",
					Link: "https://www.cdc.gov/tb/topic/basics/tbprevention.htm",
				},
			},
			Symptoms: []Symptom{
				{
					ID:     "S1",
					Name:   "Fever for two weeks or more",
					Weight: 0.513,
				},
				{
					ID:     "S2",
					Name:   "Coughing blood",
					Weight: 0.475,
				},
				{
					ID:     "S3",
					Name:   "Sputum mixed with blood",
					Weight: 0.519,
				},
				{
					ID:     "S4",
					Name:   "Night sweats",
					Weight: 0.514,
				},
				{
					ID:     "S5",
					Name:   "Chest pain",
					Weight: 0.494,
				},
				{
					ID:     "S6",
					Name:   "Back pain in certain parts",
					Weight: 0.511,
				},
				{
					ID:     "S7",
					Name:   "Shortness of breath",
					Weight: 0.487,
				},
				{
					ID:     "S8",
					Name:   "Weight loss",
					Weight: 0.521,
				},
				{
					ID:     "S9",
					Name:   "Body feels tired",
					Weight: 0.496,
				},
				{
					ID:     "S10",
					Name:   "Lumps that appear around the armpits and neck",
					Weight: 0.484,
				},
				{
					ID:     "S11",
					Name:   "Cough and phlegm continuously for two weeks to four weeks",
					Weight: 0.493,
				},
				{
					ID:     "S12",
					Name:   "Swollen lymph nodes",
					Weight: 0.478,
				},
				{
					ID:     "S13",
					Name:   "Loss of apetite",
					Weight: 0.488,
				},
			},
		},
	}

	return diseases
}

package tuneps

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

// ************************ Appel d'offres *******************************************
// AOGETlist
func REQAOGetList() (TunepsAOWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}

	//Prépare les données à envoyer dans la requête HTTP sous la forme d'une chaîne JSON
	var data = strings.NewReader(`{"listSort":[],"dataSearch":[{"key":"publicYn","value":"Y","specificSearch":"="}],"listCol":[],"pagination":{"offSet":0,"limit":90000},"sort":{"nameCol":"publicDt","direction":"desc nulls last"}}`)

	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("POST", "https://www.tuneps.tn/api/portail/bid/master/data", data)
	if err != nil {
		log.Fatal(err)
	}

	//Plusieurs en-têtes HTTP sont configurés pour simuler une requête provenant d'un navigateur.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Origin", "https://www.tuneps.tn")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOWebList pour stocker les données désérialisées JSON.
	aoList := TunepsAOWebList{}
	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoList.
	err = json.Unmarshal(bodyText, &aoList)

	if err != nil {
		return TunepsAOWebList{}, err
	}

	return aoList, nil
}

func REQAOGetInfoList() (TunepsAOInfoWebList, error) {

	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/bid/master/60079", nil)
	if err != nil {
		log.Fatal(err)
	}

	//Plusieurs en-têtes HTTP sont configurés pour simuler une requête provenant d'un navigateur.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/60079/20240206450")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOInfoWebList pour stocker les données désérialisées JSON.
	aoInfoList := TunepsAOInfoWebList{}
	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoInfoList.
	err = json.Unmarshal(bodyText, &aoInfoList)

	if err != nil {
		return TunepsAOInfoWebList{}, err
	}

	return aoInfoList, nil
}

// AOGetLot

func REQAOGetLotList() (TunepsAOLotWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/vBidCls/lot?bidNo=20240206430", nil)
	if err != nil {
		log.Fatal(err)
	}
	//Plusieurs en-têtes HTTP sont configurés pour simuler une requête provenant d'un navigateur.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/60051/20240206430")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOInfoWebList pour stocker les données désérialisées JSON.
	aoLotList := TunepsAOLotWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoLotList.
	err = json.Unmarshal(bodyText, &aoLotList)

	if err != nil {
		return TunepsAOLotWebList{}, err
	}

	return aoLotList, nil

}

// AOGetCC

func REQAOGetCCList() (TunepsAOCCWebList, error) {

	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/ged//vAttachFile/getByBidNo?bidNo=20240206430", nil)
	if err != nil {
		log.Fatal(err)
	}
	//Plusieurs en-têtes HTTP sont configurés pour simuler une requête provenant d'un navigateur.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/60051/20240206430")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOInfoWebList pour stocker les données désérialisées JSON.
	aoCCList := TunepsAOCCWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoCCList.
	err = json.Unmarshal(bodyText, &aoCCList)

	if err != nil {
		return TunepsAOCCWebList{}, err
	}

	return aoCCList, nil
}

// AOGetAg
func REQAOGetAgList() (TunepsAOAgWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/bid/data/conditions/participation/20240206458", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/60091/20240206458")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOInfoWebList pour stocker les données désérialisées JSON.
	aoAgList := TunepsAOAgWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoAgList.
	err = json.Unmarshal(bodyText, &aoAgList)

	if err != nil {
		return TunepsAOAgWebList{}, err
	}

	return aoAgList, nil
}

// AOGetProduit
func REQAOGetProdList() (TunepsAOPRWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/bid/data/product/20240206551", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/60189/20240206551")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOInfoWebList pour stocker les données désérialisées JSON.
	aoPRList := TunepsAOPRWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoPRList.
	err = json.Unmarshal(bodyText, &aoPRList)

	if err != nil {
		return TunepsAOPRWebList{}, err
	}

	return aoPRList, nil
}

//************* BOUTON List soumissionaire(s) retenu(s)******************

// AOGetDonneeDeBase
func REQAOGetDNBList() (TunepsAODNBWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/bid/master/24494", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/24494/20220600003/listSoumissionaires")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAODNBWebList pour stocker les données désérialisées JSON.
	aoDNBList := TunepsAODNBWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoDNBList.
	err = json.Unmarshal(bodyText, &aoDNBList)

	if err != nil {
		return TunepsAODNBWebList{}, err
	}

	return aoDNBList, nil
}

// AOGetAttributionDuMarché
func REQAOGetATTMList() (TunepsAOATTMWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	var data = strings.NewReader(`{"listSort":[],"dataSearch":[],"listCol":[]}`)
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("POST", "https://www.tuneps.tn/api/portail/bid/listWinners?bidNo=20220600003&bidModSeq=00", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Origin", "https://www.tuneps.tn")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/24494/20220600003/listSoumissionaires")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAOATTMWebList pour stocker les données désérialisées JSON.
	aoATTMList := TunepsAOATTMWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoATTMList.
	err = json.Unmarshal(bodyText, &aoATTMList)

	if err != nil {
		return TunepsAOATTMWebList{}, err
	}

	return aoATTMList, nil

}

// ************************bOUTON RESULTAT*************************************
//AODonnéesdebase

func REQAOGetRDNBList() (TunepsAORDNBWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/bid/master/24494", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/24494/20220600003/resultatouverture")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAORDNBWebList pour stocker les données désérialisées JSON.
	aoRDNBList := TunepsAORDNBWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoRDNBList.
	err = json.Unmarshal(bodyText, &aoRDNBList)

	if err != nil {
		return TunepsAORDNBWebList{}, err
	}

	return aoRDNBList, nil

}

func REQAOGetRLotList() (TunepsAORLotWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	var data = strings.NewReader(`{"listSort":[],"dataSearch":[],"listCol":[]}`)
	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("POST", "https://www.tuneps.tn/api/portail/bid/listLotInfructueux?bidNo=20220600003&bidModSeq=00", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Origin", "https://www.tuneps.tn")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/24494/20220600003/resultatouverture")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAORLotWebList pour stocker les données désérialisées JSON.
	aoRLotList := TunepsAORLotWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoRLotList.
	err = json.Unmarshal(bodyText, &aoRLotList)

	if err != nil {
		return TunepsAORLotWebList{}, err
	}

	return aoRLotList, nil
}

// Soumissionaires
func REQAOGetRSList() (TunepsAORSWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}

	var data = strings.NewReader(`{"listSort":[],"dataSearch":[],"listCol":[]}`)

	req, err := http.NewRequest("POST", "https://www.tuneps.tn/api/portail/ranking/rankLot/data?bidNo=20220600003&bidModSeq=00", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Origin", "https://www.tuneps.tn")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/offres/details/24494/20220600003/resultatouverture")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Déclarez une variable de type TunepsAORSWebList pour stocker les données désérialisées JSON.
	aoRSList := TunepsAORSWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoRSList.
	err = json.Unmarshal(bodyText, &aoRSList)

	if err != nil {
		return TunepsAORSWebList{}, err
	}

	return aoRSList, nil

}

/// ************************ Shopping Mall *******************************************

// ShoppingMall
func REQAOGetSMList() (TunepsSMWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}
	//Prépare les données à envoyer dans la requête HTTP sous la forme d'une chaîne JSON
	var data = strings.NewReader(`{"listSort":[],"dataSearch":[],"listCol":[],"pagination":{"offSet":0,"limit":900000},"sort":{"nameCol":"publicDt","direction":"DESC"}}`)

	// Crée une nouvelle requête HTTP POST en spécifiant l'URL cible et les données à envoyer. Si une erreur survient, le programme s'arrête en utilisant log.Fatal().
	req, err := http.NewRequest("POST", "https://www.tuneps.tn/api/portail/vSpShopMaster/data", data)
	if err != nil {
		log.Fatal(err)
	}
	//Plusieurs en-têtes HTTP sont configurés pour simuler une requête provenant d'un navigateur.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Origin", "https://www.tuneps.tn")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/consultations")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Déclarez une variable de type TunepsSMWebList pour stocker les données désérialisées JSON.
	aoSMList := TunepsSMWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoSMList.
	err = json.Unmarshal(bodyText, &aoSMList)

	if err != nil {
		return TunepsSMWebList{}, err
	}

	return aoSMList, nil

}

func REQAOGetSMInfoList() (TunepsSMInfoWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/vSpShopMaster/141839", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/consultations/consultationdetails/141839/S20240203952")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Déclarez une variable de type TunepsSMInfoWebList pour stocker les données désérialisées JSON.
	aoSMInfoList := TunepsSMInfoWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoSMInfoList.
	err = json.Unmarshal(bodyText, &aoSMInfoList)

	if err != nil {
		return TunepsSMInfoWebList{}, err
	}

	return aoSMInfoList, nil
}

// lot
func REQAOGetSMLotList() (TunepsSMLotWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/portail/consultation/data/shop/item?shopNo=S20240203952&shopModeSeq=00", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/consultations/consultationdetails/141839/S20240203952")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Déclarez une variable de type TunepsSMLotWebList pour stocker les données désérialisées JSON.
	aoSMLotList := TunepsSMLotWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoSMLotList.
	err = json.Unmarshal(bodyText, &aoSMLotList)

	if err != nil {
		return TunepsSMLotWebList{}, err
	}

	return aoSMLotList, nil
}

////Documents de la consultation

func REQAOGetSMDocList() (TunepsSMDocWebList, error) {
	// Crée une nouvelle instance de client HTTP. Cela sera utilisé pour effectuer des requêtes HTTP.
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.tuneps.tn/api/ged/spShopAttachFile/getByShopNo?shopNo=S20240203952", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "cookiesession1=678A3E2935064EC37080EDFDC12828AE")
	req.Header.Set("Referer", "https://www.tuneps.tn/portail/consultations/consultationdetails/141839/S20240203952")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//Exécute la requête HTTP en utilisant le client HTTP. Si une erreur survient, le programme s'arrête.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Reporte la fermeture du corps de la réponse HTTP jusqu'à la fin de la fonction, garantissant que le corps sera fermé même en cas d'erreur.
	defer resp.Body.Close()

	// Lit le corps de la réponse HTTP. Si une erreur survient, le programme s'arrête.
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Déclarez une variable de type TunepsSMDocWebList pour stocker les données désérialisées JSON.
	aoSMDocList := TunepsSMDocWebList{}

	// Utilisez json.Unmarshal pour désérialiser les données JSON stockées dans bodyText
	// et les mapper sur la structure aoSMDocList.
	err = json.Unmarshal(bodyText, &aoSMDocList)

	if err != nil {
		return TunepsSMDocWebList{}, err
	}

	return aoSMDocList, nil
}

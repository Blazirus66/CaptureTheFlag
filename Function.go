package main

import (
"fmt"
"io/ioutil"
"net/http"
"strings"
)

func DownloadFile(fileName string, key string, port int, keytype string) {
	url := fmt.Sprintf("http://34.77.36.161:%d", port)

	// Créer un objet http.Client pour gérer les requêtes HTTP
	client := &http.Client{}

	// Créer une requête HTTP avec la méthode POST et les données nécessaires
	req, err := http.NewRequest("POST", url, strings.NewReader(keytype+"="+key))
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return
	}

	// Ajouter l'entête Content-Type pour spécifier le format des données envoyées
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Envoyer la requête HTTP et obtenir la réponse
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête:", err)
		return
	}
	defer res.Body.Close()

	// Vérifier le code de statut de la réponse
	if res.StatusCode != 200 {
		fmt.Println("Erreur: la clé secrète est incorrecte ou le port n'est pas correct")
		return
	}

	// Lire le corps de la réponse et le stocker dans un tableau d'octets
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du corps de la réponse:", err)
		return
	}

	// Écrire le fichier téléchargé sur le disque
	err = ioutil.WriteFile("fileName.txt", body, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier:", err)
		return
	}

	fmt.Println("Fichier téléchargé avec succès")
}

func findSecretKey(port int, ch chan int) {
	url := fmt.Sprintf("http://34.77.36.161:%d", port)

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}

		fmt.Println("Clé secrète trouvée sur le port", port)
		// Utilisez la méthode Split pour décomposer la chaîne de caractères en un tableau
		// de sous-chaînes en utilisant ":" comme séparateur
		keyArray := strings.Split(string(body), ": ")

		// La clé secrète se trouve à l'index 1 du tableau
		realKey := keyArray[1]

		fmt.Println(realKey)
		ch <- port
		return
	} else {
		fmt.Println("Port", port, "ne contient pas la clé secrète")
	}
}

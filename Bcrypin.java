public class Bcrypin {

	public static void main(String[] args) throws java.security.NoSuchAlgorithmException {

		String password = "password";
		final long startTime = System.currentTimeMillis();
		String hashedPassword = BCrypt.hashpw(password, BCrypt.gensalt(14));
		final long endTime = System.currentTimeMillis();
		boolean matched = BCrypt.checkpw(password, hashedPassword);

		System.out.println(hashedPassword);
		System.out.println(matched);
		System.out.println("Total execution time: " + (endTime - startTime));

	}
	
}
